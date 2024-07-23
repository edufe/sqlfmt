select
  file_id,
  mti,
  Case mti when '1240' then 'Presentment'
           when '1442' then 'Chargeback'
           when '1644' then 'Administrative'
           when '1740' then 'FeeCollection'
           else 'Other' end as MTI_Description,
  de24,
  Case de24 when '200' then 'First Presentment'
            when '205' then 'Second Presentment (Full)'
            when '282' then 'Second Presentment (Partial)'
            when '450' then 'First Chargeback (Full)'
            when '451' then 'Arbitration Chargeback (Full)'
            when '453' then 'First Chargeback (Partial)'
            when '454' then 'Arbitration Chargeback (Partial)'
            when '603' then 'Retrieval Request'
            when '640' then 'Currency Update'
            when '680' then 'File Currency Summary'
            when '685' then 'Financial Position Detail'
            when '688' then 'Settlement Position Detail'
            when '691' then 'Message Exception'
            when '693' then 'Text Message'
            when '695' then 'File Trailer'
            when '696' then 'Financial Detail Addendum'
            when '697' then 'File Header'
            when '699' then 'File Reject'
            when '700' then 'Fee Collection (Member-generated)'
            when '780' then 'Fee Collection Return'
            when '781' then 'Fee Collection Resubmission'
            when '782' then 'Fee Collection Arbitration Return'
            when '783' then 'Fee Collection (Clearing System-generated)'
            when '790' then 'Fee Collection (Funds Transfer)'
            when '791' then 'Fee Collection (Funds Transfer Backout)'
            else 'Other'end as TRXFUNCTION,
  de3,
  Case substr(de3,1,2)
        -- /*Cardholder Account Debits*/
        when '00' then 'Purcahse (Good and Services)'
        when '01' then 'ATM Cash Withdrawal'
        when '12' then 'Cash Disbursement'
        when '17' then 'Convenience Check (available only in Collection Only messages)'
        when '18' then 'Unique Transaction (requires unique MCC)'
        when '19' then 'Fee Collection (Credit to Transaction Originator)'
         -- /*Cardholder Account Credits*/
        when '20' then 'Credit (Purchase Return)'
        when '28' then 'Payment Transaction'
        when '29' then 'Fee Collection (Debit to Transaction Originator)'
        -- /*Transfer Services*/
        when '50' then 'Payment/Balance Transfer (Available Only in Collection Only Messages)'
        else 'Other' end as CHTRXTYPE,
  count(*) cant
from records
where file_id in (
  select file_id
  from files
  where processed_date = '2')
group by mti,
  file_id,
  de24,
  de3
order by 1,2

